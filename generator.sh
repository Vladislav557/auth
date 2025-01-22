#!/bin/bash
set -euo pipefail

NAME="jwt"
mkdir "$NAME"

PRIVATE_PEM="./$NAME/private.pem"
PUBLIC_PEM="./$NAME/public.pem"

ssh-keygen -t rsa -b 2048 -m PEM -f "$PRIVATE_PEM" -q -N ""
openssl rsa -in "$PRIVATE_PEM" -pubout -outform PEM -out "$PUBLIC_PEM" 2>/dev/null

rm "$PRIVATE_PEM".pub