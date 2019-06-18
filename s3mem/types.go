/*
################################################################################
# Copyright 2019 IBM Corp. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
################################################################################
*/

package s3mem

import (
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Buckets struct {
	Buckets map[string]*Bucket
	Mux     sync.Mutex
}

//(bucket/key/version)
type Bucket struct {
	Bucket                  *s3.Bucket
	AccessControlPolicy     *s3.AccessControlPolicy
	MFA                     *string
	VersioningConfiguration *s3.VersioningConfiguration
	Objects                 map[string]*VersionedObjects
	Mux                     sync.Mutex
}

type VersionedObjects struct {
	VersionedObjects []*Object
}

type Object struct {
	Object              *s3.Object
	AccessControlPolicy *s3.AccessControlPolicy
	DeletedObject       *s3.DeletedObject
	Content             []byte
}

type Users struct {
	Users map[string]*User
}

type User struct {
	CanonicalID string
	Email       string
}
