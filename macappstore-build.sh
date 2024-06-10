#!/bin/bash

APP_CERTIFICATE="3rd Party Mac Developer Application: Andrew Arrow (755R3VBZ84)"
PKG_CERTIFICATE="3rd Party Mac Developer Installer: Andrew Arrow (755R3VBZ84)"

cp ~/Documents/store.provisionprofile epoch.app/Contents

codesign --timestamp --options=runtime -s "$APP_CERTIFICATE" -v --entitlements ./build/darwin/entitlements.plist epoch.app

productbuild --sign "$PKG_CERTIFICATE" --component epoch.app /Applications ./epoch.pkg
