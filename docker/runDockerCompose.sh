#!/bin/bash

# Note that this is probably not *ideal* as opening terminal tabs programmatically depends on the OS -
# I could probably have a test in here to check which OS you are using and open using that and have that variable passed in as a flag to runDockerCompose.sh. This is important as Docker is supposed to be a sandbox (write once run anywhere).
# Not going to do that out of laziness, but FYI. This assumes you are using a macOS.

docker rmi -f $(docker images -q)

docker-compose up -d --build

# if system == mac then ....
osascript <<END 
tell application "Terminal"
    do script "cd \"`pwd`\"; docker-compose exec instance1 sh;"
end tell
END
osascript <<END 
tell application "Terminal"
    do script "cd \"`pwd`\"; docker-compose exec instance2 sh;"
end tell
END
# if system == linux then...etc

# Scratch

# DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
# ./runInstance.sh $DIR
# docker-compose exec instance1 sh 
# echo Welcome to instance1
# argArr=($DIR 'instance1')
# echo "value of args ${argArr[1]}"
# open -a Terminal.app runInstance.sh $argArr
# open -a Terminal.app runInstance.sh "instance1" "$DIR"
# open -a Terminal.app runInstance.sh "instance2" "$DIR"
# osascript -e 'tell app "Terminal"
#     do script "docker-compose exec instance1 sh; echo Welcome to instance1"
# end tell'
# osascript -e 'tell app "Terminal"
#     do script "docker-compose exec instance2 sh; echo Welcome to instance2"
# end tell'
# docker-compose exec instance2 sh 
# echo Welcome to instance2