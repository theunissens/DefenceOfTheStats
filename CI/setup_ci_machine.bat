REM The first thing we need to do is install ruby, as we will be using it for the rest of the setup of IMQS environment
mkdir c:\temp
copy "T:\Software Installation Files\CI\Tools\rubyinstaller-2.1.4-x64.exe" c:\temp
c:
REM First we need to check if ruby is installed, before running the installer
ruby --version
cd \temp
rubyinstaller-2.1.4-x64.exe /silent /loadinf="T:\Software Installation Files\CI\Tools\ruby.inf"
REM After installing ruby, we want to make sure we can run ruby commands without needing to provide extentions (.rb)
assoc .rb=RubyScript
ftype RubyScript=C:\Ruby22-x64\bin\ruby.exe %1 %*
setx /M PATHEXT %PATHEXT%;.ruby
