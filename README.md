# machodump
Golang tool to dump useful information from a Mach-O binary. Read more on our [blog post](https://redmaple.tech/blogs/macho-files/).

Uses the library [Go-Macho](https://github.com/blacktop/go-macho) for low-level parsing.

There is a handy [Mach-o file template](https://www.sweetscape.com/010editor/repository/files/MachO.bt) for 010 Editor.

# Build
Just run:

```
make
```

# Example

```
./machodump -i testfiles/Chess
2020/11/24 14:20:29 Parsing file "testfiles/Chess".
File Details:
        Magic: 64-bit MachO
        Type: Exec
        CPU: Amd64, x86_64
        Commands: 31 (Size: 4328)
        Flags: NoUndefs, DyldLink, TwoLevel, BindsToWeak, PIE
        UUID: 18455A71-F835-3D0F-8F7C-215BF86BC7AF
File imports 15 libraries:
        0: "/System/Library/Frameworks/SystemConfiguration.framework/Versions/A/SystemConfiguration"
        1: "/System/Library/Frameworks/CoreFoundation.framework/Versions/A/CoreFoundation"
        2: "/System/Library/Frameworks/AppKit.framework/Versions/C/AppKit"
        3: "/System/Library/Frameworks/GameKit.framework/Versions/A/GameKit"
        4: "/System/Library/Frameworks/Cocoa.framework/Versions/A/Cocoa"
        5: "/System/Library/Frameworks/OpenGL.framework/Versions/A/OpenGL"
        6: "/System/Library/Frameworks/Carbon.framework/Versions/A/Carbon"
        7: "/System/Library/Frameworks/Foundation.framework/Versions/C/Foundation"
        8: "/usr/lib/libobjc.A.dylib"
        9: "/usr/lib/libc++.1.dylib"
        10: "/usr/lib/libSystem.B.dylib"
        11: "/System/Library/Frameworks/ApplicationServices.framework/Versions/A/ApplicationServices"
        12: "/System/Library/Frameworks/CoreGraphics.framework/Versions/A/CoreGraphics"
        13: "/System/Library/Frameworks/CoreServices.framework/Versions/A/CoreServices"
        14: "/System/Library/Frameworks/ImageIO.framework/Versions/A/ImageIO"
File has 31 load commands. Interesting commands:
        Load 11 (LC_SOURCE_VERSION): 369.0.0.0.0
Binary has 1 Code Directory:
        CodeDirectory 0:
                Ident: "com.apple.Chess"
                CD Hash: 9a95a73ca9b45ad1f0a603b0045c8baf256c289e959d491082acb063e81d30c9
                Code slots: 68
                Special slots: 5
                        Special Slot   5 Entitlements Blob:     560ea26d0d71b1927a954f02955932a7686c942730ecbf08736141c2e3893f00
                        Special Slot   4 Application Specific:  Not Bound
                        Special Slot   3 Resource Directory:    a5dc7f7078d841ff1783af7e5940e078e67952916c44eb5598092c133a6585e3
                        Special Slot   2 Requirements Blob:     0a7f89d52a71b16993861945aedc98a50019d9c84bf95e3cc4a88c37a68720ba
                        Special Slot   1 Bound Info.plist:      f7c596d135a14757e7444eacc9680924f52882b59d480fe4db5c40308e88cffa
CMS Signature has 3 certificates:
        CN: "Apple Code Signing Certification Authority"
        CN: "Apple Root CA"
        CN: "Software Signing"
Binary has 1 requirement:
        Requirement 0 (Designated Requirement): identifier "com.apple.Chess" and anchor apple
Binary has 5 boolean entitlements:
        com.apple.developer.game-center: true
        com.apple.security.app-sandbox: true
        com.apple.security.device.microphone: true
        com.apple.security.files.user-selected.read-write: true
        com.apple.security.network.client: true
Binary has 1 string array entitlement:
        0 com.apple.private.tcc.allow: ["kTCCServiceMicrophone"]
2020/11/24 14:20:29 Fin.
```
