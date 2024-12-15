#!/usr/bin/bash
# Author Prince kumar 
# this scritp is going to be insallation of all the addtional script in bash...
# date 28 nov 2024
# first check for the os 
command -v go 2>&1 > /dev/null || { echo -e  "${g}+++++${y}Installing golang ${g}+++++" ; apt-get install golang -y ; }
check_os=$(uname -o)
if [[ $check_os == "GNU/Linux" ]];then 
    # install the go 
    # now make the binary and insatall the program to the path
    go build -o style 
    # Now move the binary to the bin/bash 
    go install 
elif [[ $check_os == "Android" ]]; then 
    # install the go for android ..
    go build -o style 
    go install
    mv style /data/data/com.termux/files/usr/bin
else 
    echo ""
fi



