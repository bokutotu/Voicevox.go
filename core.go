package voicevox

/*
#cgo LDFLAGS: -Wl,-rpath,${SRCDIR}/bin
#cgo LDFLAGS: -L${SRCDIR}/bin 
#cgo LDFLAGS: -lvoicevox_core
#include "./bin/voicevox_core.h"
*/
import "C"

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
    if result == 0 {
        return output_wav, nil
    }
    return output_wav, VoicevoxError(result)
}
