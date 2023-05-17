package voicevox

/*
#cgo LDFLAGS: -Wl,-rpath,${SRCDIR}/bin
#cgo LDFLAGS: -L${SRCDIR}/bin 
#cgo LDFLAGS: -lvoicevox_core
#include "./bin/voicevox_core.h"
*/
import "C"

import (
    "fmt"
    // "unsafe"
)

type VoicevoxError C.int

const (
    LOADED_OPENJTALK_DICT_ERROR VoicevoxError = C.VOICEVOX_RESULT_NOT_LOADED_OPENJTALK_DICT_ERROR
    LOAD_MODEL_ERROR VoicevoxError = C.VOICEVOX_RESULT_LOAD_MODEL_ERROR
    GET_SUPPORTED_DEVICES_ERROR VoicevoxError = C.VOICEVOX_RESULT_GET_SUPPORTED_DEVICES_ERROR
    GPU_SUPPORT_ERROR VoicevoxError = C.VOICEVOX_RESULT_GPU_SUPPORT_ERROR
    LOAD_METAS_ERROR VoicevoxError = C.VOICEVOX_RESULT_LOAD_METAS_ERROR
    UNIINITIALIZED_STATUS_ERROR VoicevoxError = C.VOICEVOX_RESULT_UNINITIALIZED_STATUS_ERROR
    INVALID_SPEAKER_ID_ERROR VoicevoxError = C.VOICEVOX_RESULT_INVALID_SPEAKER_ID_ERROR
    INVALID_MODEL_INDEX_ERROR VoicevoxError = C.VOICEVOX_RESULT_INVALID_MODEL_INDEX_ERROR
    INFERENCE_ERROR VoicevoxError = C.VOICEVOX_RESULT_INFERENCE_ERROR
    EXTRACT_FULL_CONTEXT_LABEL_ERROR VoicevoxError = C.VOICEVOX_RESULT_EXTRACT_FULL_CONTEXT_LABEL_ERROR
    INVALID_UTF8_INPUT_ERROR VoicevoxError = C.VOICEVOX_RESULT_INVALID_UTF8_INPUT_ERROR
    PARSE_KANA_ERROR VoicevoxError = C.VOICEVOX_RESULT_PARSE_KANA_ERROR
    INVALID_AUDIO_QUERY_ERROR VoicevoxError = C.VOICEVOX_RESULT_INVALID_AUDIO_QUERY_ERROR
)

func from_int(i int) VoicevoxError {
    return VoicevoxError(C.int(i))
}

func (e VoicevoxError) Error() string {
    str := C.voicevox_error_result_to_message(C.int(e))
    go_str := C.GoString(str)
    return fmt.Sprintf("VoicevoxError: %s", go_str)
}
