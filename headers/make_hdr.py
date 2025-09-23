import sys
if len(sys.argv) != 2:
    print("please enter 1 of")
    print("jpg, jpeg, png, gif")
    sys.exit()


newFileBytes = []
ext = sys.argv[1]
match(ext):
    case "jpg":
        newFileBytes = [255, 216, 255, 224]
    case "jpeg":
        newFileBytes = [255, 216, 255, 224]
    case "png":
        newFileBytes = [137, 80, 78, 71, 13, 10, 26, 10]
    case "gif":
        newFileBytes = [71, 73, 70, 56, 57, 97]

newFileBytesArray = bytearray(newFileBytes)
with open("test." + ext, "wb") as fd:
    fd.write(newFileBytesArray)
