// Copyright (c) 2015 Pagoda Box Inc
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v.
// 2.0. If a copy of the MPL was not distributed with this file, You can obtain one
// at http://mozilla.org/MPL/2.0/.
//

package config

import (
	"os"
)

// ParseVMfile
func ParseVMfile() *VMfileConfig {

	//
	vmfile := &VMfileConfig{}
	vmfilePath := AppDir + "/.vmfile"

	// if a .vmfile doesn't exist - create it
	if _, err := os.Stat(vmfilePath); err != nil {

		vmfile.Deployed = false
		vmfile.Mode = "foreground"
		vmfile.Status = "not created"
		vmfile.Suspendable = true

		writeVMfile()

		// if a .vmfile does exists - parse it
	} else {
		if err := ParseConfig(vmfilePath, vmfile); err != nil {
			Fatal("[config/vmfile] ParseConfig() failed", err.Error())
		}
	}

	return vmfile
}

//
func (c *VMfileConfig) HasDeployed() bool {
	if err := ParseConfig(AppDir+"/.vmfile", c); err != nil {
		Fatal("[config/vmfile] ParseConfig() failed", err.Error())
	}

	return c.Deployed
}

//
func (c *VMfileConfig) IsMode(mode string) bool {
	if err := ParseConfig(AppDir+"/.vmfile", c); err != nil {
		Fatal("[config/vmfile] ParseConfig() failed", err.Error())
	}

	return c.Mode == mode
}

//
func (c *VMfileConfig) IsStatus(status string) bool {
	if err := ParseConfig(AppDir+"/.vmfile", c); err != nil {
		Fatal("[config/vmfile] ParseConfig() failed", err.Error())
	}

	return c.Status == status
}

//
func (c *VMfileConfig) IsSuspendable() bool {
	if err := ParseConfig(AppDir+"/.vmfile", c); err != nil {
		Fatal("[config/vmfile] ParseConfig() failed", err.Error())
	}

	return c.Suspendable
}

//
func (c *VMfileConfig) DeployedIs(deployed bool) {
	c.Deployed = deployed
	writeVMfile()
}

//
func (c *VMfileConfig) ModeIs(mode string) {
	c.Mode = mode
	writeVMfile()
}

//
func (c *VMfileConfig) StatusIs(status string) {
	c.Status = status
	writeVMfile()
}

//
func (c *VMfileConfig) SuspendableIs(suspendable bool) {
	c.Suspendable = suspendable
	writeVMfile()
}

//
func (c *VMfileConfig) UUIDIs(uuid string) {
	c.UUID = uuid
	writeVMfile()
}

// writeVMfile
func writeVMfile() {
	if err := writeConfig(AppDir+"/.vmfile", VMfile); err != nil {
		Fatal("[config/vmfile] writeConfig() failed", err.Error())
	}
}
