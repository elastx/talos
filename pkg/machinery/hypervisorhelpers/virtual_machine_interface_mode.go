// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package hypervisorhelpers

// VirtualMachineInterfaceMode is how a virtual machine interface is attached to a host link.
type VirtualMachineInterfaceMode int

// VirtualMachineInterfaceMode constants.
//
//structprotogen:gen_enum
const (
	VirtualMachineInterfaceModeUnknown VirtualMachineInterfaceMode = iota // unknown
	VirtualMachineInterfaceModeDirect                                     // direct
	VirtualMachineInterfaceModeBridge                                     // bridge
)
