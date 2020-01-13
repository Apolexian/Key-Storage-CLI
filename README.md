Key Storage Secret CLI
------------------------

Meant to be a better way to store API keys than a plaintext
file but does not provide full security against attacks. Mainly for 
convenience of storage and *some* security. 

Code Style
---------------------------------
Trying to stick to best go practises as much as possible:
https://github.com/golang/go/wiki/CodeReviewComments 

Status
-------------------------------------
Currently done with Encryption/Decryption algorithms. <br>
Added tests for encryption and decryption. <br>
Added a in memory storage version for initial testing purposes. <br>
Added Makefile for project. <br>
Added initial file backed version. <br>
Added test files for storage. <br>
Completed the base CLI shell. <br>
Added logger.go to handle logging more gracefully. <br>

