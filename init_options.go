package voicevox

/*
#cgo LDFLAGS: -Wl,-rpath,${SRCDIR}/bin
#cgo LDFLAGS: -L${SRCDIR}/bin 
#cgo LDFLAGS: -lvoicevox_core
#include "./bin/voicevox_core.h"
*/
import "C"

type InitOptions struct {
    _options C.VoicevoxInitializeOptions
}

func DefaultInitOtions() InitOptions {
    options := InitOptions{_options: C.voicevox_make_default_initialize_options() }
    return options
}

func (options *InitOptions) SetLoadAllModels(flag bool) {
    options._options.load_all_models = C.bool(flag)
}

func (options *InitOptions) SetOpenJTalkDictDir(path string) {
    options._options.open_jtalk_dict_dir = C.CString(path)
}

func (options *InitOptions) SetAccelerationMode(mode Acceleration) {
    options._options.acceleration_mode = C.int(mode)
}

func (options *InitOptions) SetNumThreads(num int) {
    options._options.cpu_num_threads = C.ushort(num)
}
