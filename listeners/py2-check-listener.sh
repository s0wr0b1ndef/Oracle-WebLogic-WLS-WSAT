#!/bin/sh

echo "[+] Installing requests dependency"
pip install -U requests

echo "[+] Starting listener on port 4444"
python -m SimpleHTTPServer 4444
