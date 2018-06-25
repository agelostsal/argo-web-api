/*
 * Copyright (c) 2018 GRNET S.A.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the
 * License. You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an "AS
 * IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language
 * governing permissions and limitations under the License.
 *
 * The views and conclusions contained in the software and
 * documentation are those of the authors and should not be
 * interpreted as representing official policies, either expressed
 * or implied, of GRNET S.A.
 *
 */

package thresholdsProfiles

import (
	"github.com/ARGOeu/argo-web-api/respond"
	"github.com/gorilla/mux"
)

// HandleSubrouter uses the subrouter for a specific calls and creates a tree of sorts
// handling each route with a different subrouter
func HandleSubrouter(s *mux.Router, confhandler *respond.ConfHandler) {

	s = respond.PrepAppRoutes(s, confhandler, appRoutesV2)

}

var appRoutesV2 = []respond.AppRoutes{
	{"thresholdsProfiles.list", "GET", "/thresholds_profiles", List},
	{"thresholdsProfiles.get", "GET", "/thresholds_profiles/{ID}", ListOne},
	{"thresholdsProfiles.create", "POST", "/thresholds_profiles", Create},
	{"thresholdsProfiles.update", "PUT", "/thresholds_profiles/{ID}", Update},
	{"thresholdsProfiles.delete", "DELETE", "/thresholds_profiles/{ID}", Delete},
	{"thresholdsProfiles.options", "OPTIONS", "/thresholds_profiles", Options},
	{"thresholdsProfiles.options", "OPTIONS", "/thresholds_profiles/{ID}", Options},
}
