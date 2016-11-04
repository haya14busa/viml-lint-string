
let _ = "double quote"

let _ = "ok if it contains 'single quote'"

let _ = map([], 'ok "double quote" in map expr')

let _ = "double quote 2"

" Single quote literal should not be prohibited
echo ''
echo 'foo'

" Double quote literal should not be prohibited when it contains escape sequence such as line break
echo "\001"
echo "\xff"
echo "\uffff"
echo "\b"
echo "\e"
echo "\f"
echo "\n"
echo "\t"
echo "\\"
echo "\""
echo "\<xxx>"

" Double quote literal should not be prohibited too
echo "'"

