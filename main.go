package main

// #cgo CFLAGS: -I${SRCDIR}/libs/androidfw
// #cgo LDFLAGS: -lstdc++ -L${SRCDIR}/libs/androidfw -landroidfw
// #include "androidfw/ApkAssets.h"
// #include "androidfw/Asset.h"
// #include "androidfw/AssetDir.h"
// #include "androidfw/AssetManager.h"
// #include "androidfw/AssetManager2.h"
// #include "androidfw/AttributeFinder.h"
// #include "androidfw/AttributeResolution.h"
// #include "androidfw/BackupHelpers.h"
// #include "androidfw/ByteBucketArray.h"
// #include "androidfw/Chunk.h"
// #include "androidfw/ConfigDescription.h"
// #include "androidfw/CursorWindow.h"
// #include "androidfw/DisplayEventDispatcher.h"
// #include "androidfw/Idmap.h"
// #include "androidfw/LoadedArsc.h"
// #include "androidfw/Locale.h"
// #include "androidfw/LocaleData.h"
// #include "androidfw/misc.h"
// #include "androidfw/MutexGuard.h"
// #include "androidfw/ObbFile.h"
// #include "androidfw/PosixUtils.h"
// #include "androidfw/ResourceTypes.h"
// #include "androidfw/ResourceUtils.h"
// #include "androidfw/StreamingZipInflater.h"
// #include "androidfw/StringPiece.h"
// #include "androidfw/TypeWrappers.h"
// #include "androidfw/Util.h"
// #include "androidfw/ZipFileRO.h"
// #include "androidfw/ZipUtils.h"
import "C"

func main() {
	C.bar()
}
