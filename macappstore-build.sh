#!/bin/bash

APP_CERTIFICATE="3rd Party Mac Developer Application: YOUR NAME (CODE)"
PKG_CERTIFICATE="3rd Party Mac Developer Installer: YOUR NAME (CODE)"
APP_NAME="YourApp"

wails build -platform darwin/universal -clean

cp ./embedded.provisionprofile "./build/bin/$APP_NAME.app/Contents"

codesign --timestamp --options=runtime -s "$APP_CERTIFICATE" -v --entitlements ./build/darwin/entitlements.plist ./build/bin/$APP_NAME.app

productbuild --sign "$PKG_CERTIFICATE" --component ./build/bin/$APP_NAME.app /Applications ./$APP_NAME.pkg
