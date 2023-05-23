package voicevox

/*
#cgo LDFLAGS: -Wl,-rpath,${SRCDIR}/bin
#cgo LDFLAGS: -L${SRCDIR}/bin
#cgo LDFLAGS: -lvoicevox_core
#include "./bin/voicevox_core.h"
*/
import "C"
import "unsafe"

func VoicevoxInitialize(options InitOptions) VoicevoxError {
    result := C.voicevox_initialize(options._options)
    return VoicevoxError(result)
}

func Tts(text string, speaker_id int, options TtsOptions) ([]byte, error) {
    output_wav_size_c := C.ulong(0)
    output_wav_ptr := (*C.uchar)(nil)
    result := C.voicevox_tts(
        C.CString(text),
        C.uint(speaker_id),
        options._options,
        &output_wav_size_c,
        &output_wav_ptr,
    )
    var output_wav []byte = make([]byte, output_wav_size_c)
    output_wav_size := int(output_wav_size_c)
    tmp := C.GoBytes(unsafe.Pointer(&output_wav_ptr), C.int(output_wav_size))
    copy(output_wav, tmp)
    if result == 0 {
        return output_wav, nil
    }
    C.voicevox_wav_free(output_wav_ptr)
    return output_wav, VoicevoxError(result)
}
