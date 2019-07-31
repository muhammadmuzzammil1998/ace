#!/usr/bin/env bash

# MIT License
#
# Copyright (c) 2019 Muhammad Muzzammil
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.

platforms=("darwin/386" "darwin/amd64" "dragonfly/amd64" "freebsd/386" "freebsd/amd64" "freebsd/arm" "linux/386" "linux/amd64" "linux/arm" "linux/arm64" "linux/ppc64" "linux/ppc64le" "linux/mips" "linux/mipsle" "linux/mips64" "linux/mips64le" "netbsd/386" "netbsd/amd64" "netbsd/arm" "openbsd/386" "openbsd/amd64" "openbsd/arm" "plan9/386" "plan9/amd64" "solaris/amd64" "windows/386" "windows/amd64")
cd ..
rm -rf bin deb
mkdir bin

pname=$(basename $(pwd))

echo "Enter details for" $pname
read -p "Version: " version
read -p "Maintainer: " maintainer
read -p "Repository URL: " url
read -p "Description: " desc
read -p "License: " License

for platform in "${platforms[@]}"; do
  package_name=$pname

  platform_split=(${platform//\// })
  GOOS=${platform_split[0]}
  GOARCH=${platform_split[1]}
  output_name=$package_name'-'$GOOS'-'$GOARCH

  if [ $GOOS = "windows" ]; then
    package_name+='.exe'
  fi

  env GOOS=$GOOS GOARCH=$GOARCH go build -o $package_name

  if [ $? -ne 0 ]; then
    echo 'Unable to build for '$GOOS'-'$GOARCH
  else
    if [ $GOOS = "linux" ]; then
      if [[ $GOARCH == "386" || $GOARCH == "amd64" ]]; then
        mkdir deb
        cd deb
        mkdir $pname
        mkdir $pname/DEBIAN
        mkdir $pname/usr
        mkdir $pname/usr/bin
        cp ../$package_name $pname/usr/bin
        echo -e "Package: $pname\nVersion: $version\nArchitecture: all\nMaintainer: $maintainer\nHomepage: $url\nDescription: $desc\nLicense: $license\n" >$pname/DEBIAN/control

        dpkg-deb --build $pname
        mv $pname.deb ../bin/$output_name'.deb'
        cd ../bin
        file $output_name'.deb' >>BIN_INFO
        cd ..
        rm -rf deb
      fi
    fi

    if [ $GOOS = "windows" ]; then
      cp $package_name bin/$output_name'.exe'
    fi

    mkdir $output_name
    cp LICENSE $output_name
    mv $package_name $output_name
    echo 'Built for '$GOOS'-'$GOARCH
    file $output_name'/'$package_name >>bin/BIN_INFO
    tar -zcvf $output_name.tar.gz $output_name >/dev/null
    echo 'Packed in '$output_name'.tar.gz'
    mv $output_name'.tar.gz' bin/
    rm -rf $output_name
  fi
done

echo -e "\n\nPacking binaries..."
tar -zcvf .dist/$pname-binaries-$version.tar.gz bin >/dev/null

rm -rf bin
