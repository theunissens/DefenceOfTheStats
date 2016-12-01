# The first thing we need to do is install ruby, as we will be using it for the rest of the setup of IMQS environment
mkdir c:\temp -ErrorAction SilentlyContinue
Copy-Item "T:\Software Installation Files\CI\Tools\rubyinstaller-2.1.4-x64.exe" c:\temp
# First makes ure ruby is not installed on this machine
Try {
    $OutputVariable = ruby --version | Out-String
    if ($OutputVariable -match "\bruby\b [0-9]") {
        $InstallRuby = $FALSE
    }
}
Catch {
    # An error ocurred, it means ruby is not installed as a program, and we should install it
    $InstallRuby = $TRUE
}

if ($InstallRuby -eq $TRUE) {
    write-host "Installing ruby..."
    Set-Location C:\temp
    $RubyInstallationProcess = 'rubyinstaller-2.1.4-x64.exe /silent /loadinf="T:\Software Installation Files\CI\Tools\ruby.inf'
    #$myarg = '-i ConnectionUsage=DSS Port=3311 ServiceName=MySQL RootPassword= ' + $rootpw
    #rubyinstaller-2.1.4-x64.exe /silent /loadinf="T:\Software Installation Files\CI\Tools\ruby.inf"  
    Start-Process $RubyInstallationProcess
} else {
    write-host "Ruby already installed"
}
