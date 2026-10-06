#!/usr/bin/env bash
# Builds the TCHES camera-ready bundle: stages only the sources needed by
# ches.tex, compiles them in isolation, and writes camera-ready/*.{zip,pdf}.
set -euo pipefail
cd "$(dirname "$0")"

NAME=tches2027-185-srivastav
OUT=camera-ready
STAGE=$(mktemp -d)/$NAME
mkdir -p "$STAGE/sections" "$OUT"

cp ches.tex references.bib iacrtrans.cls CC-by.pdf "$STAGE/"
for f in $(grep -o 'sections/[^}]*\.tex' ches.tex); do cp "$f" "$STAGE/sections/"; done

(cd "$STAGE" && latexmk -pdf -interaction=nonstopmode -halt-on-error ches.tex >/dev/null)
grep -q "undefined" "$STAGE/ches.log" && { echo "undefined refs/citations"; exit 1; }

# Keep the .bbl so the source builds without running BibTeX.
(cd "$STAGE" && latexmk -c ches.tex >/dev/null && rm -f ches.synctex.gz)
cp "$STAGE/ches.pdf" "$OUT/$NAME.pdf"
rm -f "$OUT/$NAME.zip"
(cd "$STAGE/.." && zip -qr - "$NAME" -x "$NAME/ches.pdf") > "$OUT/$NAME.zip"
unzip -l "$OUT/$NAME.zip"
