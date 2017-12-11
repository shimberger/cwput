#!/bin/bash

# Configure region and namespace
REGION="eu-central-1"
NAMESPACE="TestNamespace"

# Shorthand to make putting metrics easier
CMD="./cwput --region $REGION --namespace "${NAMESPACE}" --dimensions Server=$(hostname -s)"

# Save load average
LOAD_AVG="$(uptime | awk -F'[a-z]:' '{ print $2}' | cut -d' ' -f 2)"
$CMD --metric "Load Average" --unit "Count" --value="${LOAD_AVG}"

# Save free disk space
DISKSPACE=$(df -h | grep /dev/disk1s1 | awk '/[0-9]%/{print $(5)}' | tr -d '%')
$CMD --metric "Free Disk Space (disk1s1)" --unit "Percent" --value="${DISKSPACE}"