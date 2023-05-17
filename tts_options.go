package voicevox

/*
#cgo LDFLAGS: -Wl,-rpath,${SRCDIR}/bin
#cgo LDFLAGS: -L${SRCDIR}/bin 
#cgo LDFLAGS: -lvoicevox_core
#include "./bin/voicevox_core.h"
*/
import "C"

type TtsOptions struct {
    _options C.VoicevoxTtsOptions
}

func DefaultTtsOption() TtsOptions {
    option := TtsOptions {_options: C.voicevox_make_default_tts_options()}
    return option
}

func (option *TtsOptions) SetKana(flag bool) {
    option._options.kana = C.bool(flag)
}

func (options *TtsOptions) SetEnableInterrogativeUpspeak(flag bool) {
    options._options.enable_interrogative_upspeak = C.bool(flag)
}
