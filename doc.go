// Package ntlmv2hash computes weak, unsalted, hashes of passwords
// in the the Microsoft Windows NT LAN Manager (NTLM) protocol v2
// format.
//
// It's primary purpose is to convert passwords for use with Samba and
// its `pdbedit --set-nt-hash` tool. Using NTLMv2 hashes for anything
// else is not advised.
//
// See `cmd` for the command-line tool.
//
// Resources
//
// - https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-nlmp/b38c36ed-2804-4868-a9ff-8dd3182128e4
//
// - http://davenport.sourceforge.net/ntlm.html
//
// - https://github.com/Azure/go-ntlmssp (doesn't expose this
// functionality)
package ntlmv2hash
