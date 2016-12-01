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
    
}