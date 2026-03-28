#!/bin/sh
echo "<html><body>" > index.html
echo "<h1>System Information</h1>" >> index.html
echo "<p><strong>IP Address:</strong> $(hostname -i)</p>" >> index.html
echo "<p><strong>Hostname:</strong> $(hostname)</p>" >> index.html
echo "<p><strong>App Version:</strong> $APP_VERSION</p>" >> index.html
echo "</body></html>" >> index.html