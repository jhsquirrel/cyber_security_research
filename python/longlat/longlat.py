import sys
if len(sys.argv) != 2:
    print(sys.argv[0] + " <file containing exiftool longitude latitude>")
    print("e.g. 51 deg 30' 51.90\" N, 0 deg 5' 38.73\" W >")
    sys.exit()
with open(sys.argv[1], "r") as f:
    longlat = f.read()
# eg. 51 deg 30' 51.90" N, 0 deg 5' 38.73" W
print(longlat)
longlat = longlat.replace(" ", "")
longlat = longlat.replace("deg", chr(176))
print(longlat)
