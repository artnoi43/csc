# csc - a CLI file checksum checker
csc is a simple, stupid file checksum checker. It currently supports SHA256 (default), SHA1, and MD5 - all of which are popular hash algorithms.

## Depedencies
csc only uses the Go standard library.
## Usage
csc needs a filename, and (optionally) a checksum to compare against. Users can toggle between the hashing algorithms with `-a` command-line switch.
### Syntax

    csc -a <ALGO> -f <FILE> -c [CHECKSUM]

Examples

    csc -f downloaded.dmg -c 13d8bf76dbf5b3877c4e17946b967e2005cec8fdd2952961a1c282ed3301053f
    csc -f downloaded.dmg -a md5 -c 13e98b61102d10daf96fb335609b537c

### Omitting the comparison part and only calculating the checksum
csc will still prints the file checksum. although it exits with non-zero status. This is because omitting the checksum is equivalent to having the comparison function compare a string (the calculated checksum) to an empy string (the one you omitted).
