package libvirt

/*
#cgo pkg-config: libvirt
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <assert.h>
#include "node_device_compat.h"
#include "node_device_events_cfuncs.h"
#include "callbacks_cfuncs.h"
#include <stdint.h>

extern void nodeDeviceEventLifecycleCallback(virConnectPtr, virNodeDevicePtr, int, int, int);
void nodeDeviceEventLifecycleCallback_cgo(virConnectPtr c, virNodeDevicePtr d,
                                           int event, int detail, void *data)
{
    nodeDeviceEventLifecycleCallback(c, d, event, detail, (int)(intptr_t)data);
}

extern void nodeDeviceEventGenericCallback(virConnectPtr, virNodeDevicePtr, int);
void nodeDeviceEventGenericCallback_cgo(virConnectPtr c, virNodeDevicePtr d, void *data)
{
    nodeDeviceEventGenericCallback(c, d, (int)(intptr_t)data);
}

int virConnectNodeDeviceEventRegisterAny_cgo(virConnectPtr c,  virNodeDevicePtr d,
                                              int eventID, virConnectNodeDeviceEventGenericCallback cb,
                                              long goCallbackId) {
    void* id = (void*)goCallbackId;
#if LIBVIR_VERSION_NUMBER < 2002000
    assert(0); // Caller should have checked version
#else
    return virConnectNodeDeviceEventRegisterAny(c, d, eventID, cb, id, freeGoCallback_cgo);
#endif
}

*/
import "C"
