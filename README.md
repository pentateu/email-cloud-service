# smtp cloud service
THis component is responsible to receive email messages, encrypt and store on IPFS until is retrieved by nodes or until expiration time.

## Features:
### Receiving emails:
1) Listens for emails
1.1) Filter incomoing emails
2) Encrypt Message
3) Store message in IPFS


### Clean-up expired folders Daemon
Remove folders from local node (not from IPFS) that have expired.

## Initially Based on: github.com/flashmob/maildiranasaurus
maildiranasaurus was a great starting point.

# Mail Libraries
https://github.com/flashmob


# Encryption
https://github.com/google/tink/blob/master/docs/GOLANG-HOWTO.md
https://pkg.go.dev/filippo.io/age

## OpenPGP
Initially OpenPGP was considered to encrypt all email messages since it is the most common email encryption mechanism out there, but golang devs basicallty recommends agains OpenPGP(https://github.com/golang/go/issues/44226). The SMTP Service will use a more modern and simplified encryption mechanism.

For legacy compactibility OpenPGP must be supported by the client, but is not required in the smtp service.

### Go OpenPGP and Encryption Library
https://github.com/ProtonMail/gopenpgp



# storing messages
## MailDir
https://en.wikipedia.org/wiki/Maildir
http://www.courier-mta.org/maildir.html

MailDir is an stardard to store email messages in the filesystem using a file for each message organized in folders.



## IPFS
IPFS is used to sync the mail folders

### RFC-5322

ipld-eml is an RFC-5322 compliand IPLD object format for storing email messages
