# PaperCut Auth

A trivial example of a custom user authentication program for use with PaperCut NG or MF See [Custom User Sync Integration](http://www.papercut.com/kb/Main/CaseStudyCustomUserSyncIntegration)

This example is for illustrative purposes only. Any production solution needs to be optimised for performance.

## Setting up PaperCut

* Put the application in [install_path]/server/custom
* Go to Options > User / Group Sync
* Change the Primary Sync Source to Custom Program (Advanced)
* Set the Custom User Program to whatever you need for your Groups (Windows AD would be [install_path]/server/bin/win/UserDirAD.exe)
* Set the Custom Auth Program to [install_path]/server/custom/papercutauth.exe