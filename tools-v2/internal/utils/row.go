/*
 *  Copyright (c) 2022 NetEase Inc.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

/*
 * Project: CurveCli
 * Created Date: 2022-05-25
 * Author: chengyi (Cyber-SiKu)
 */

package cobrautil

const (
	ROW_ADDR            = "addr"
	ROW_BLOCKSIZE       = "blocksize"
	ROW_CAPACITY        = "capacity"
	ROW_CHILD_LIST      = "childList"
	ROW_CHILD_TYPE      = "childType"
	ROW_COPYSET_ID      = "copysetId"
	ROW_COPYSET_KEY     = "copysetKey"
	ROW_DUMMY_ADDR      = "dummyAddr"
	ROW_END             = "end"
	ROW_EPOCH           = "epoch"
	ROW_EXPLAIN         = "explain"
	ROW_EXTERNAL_ADDR   = "externalAddr"
	ROW_FS_ID           = "fsId"
	ROW_FS_NAME         = "fsName"
	ROW_FS_TYPE         = "fsType"
	ROW_HOSTNAME        = "hostname"
	ROW_ID              = "id"
	ROW_INTERNAL_ADDR   = "internalAddr"
	ROW_KEY             = "key"
	ROW_LEADER_PEER     = "leaderPeer"
	ROW_LEFT            = "left"
	ROW_METASERVER_ADDR = "metaserverAddr"
	ROW_MOUNT_NUM       = "mountNum"
	ROW_MOUNTPOINT      = "mountpoint"
	ROW_NAME            = "name"
	ROW_NUM             = "num"
	ROW_ONLINE_STATE    = "onlineState"
	ROW_OPERATION       = "operation"
	ROW_OWNER           = "owner"
	ROW_PARENT          = "parent"
	ROW_PARTITION_ID    = "partitionId"
	ROW_PEER_ADDR       = "peerAddr"
	ROW_PEER_ID         = "peerId"
	ROW_PEER_NUMBER     = "peerNumber"
	ROW_POOL_ID         = "poolId"
	ROW_READONLY        = "readonly"
	ROW_RESULT          = "result"
	ROW_START           = "start"
	ROW_STATE           = "state"
	ROW_STATUS          = "status"
	ROW_SUM_IN_DIR      = "sumInDir"
	ROW_TERM            = "term"
	ROW_TOTAL           = "total"
	ROW_TYPE            = "type"
	ROW_USED            = "used"
	ROW_VERSION         = "version"

	ROW_VALUE_ADD     = "add"
	ROW_VALUE_DEL     = "del"
	ROW_VALUE_DNE     = "DNE"
	ROW_VALUE_OFFLINE = "offline"
	ROW_VALUE_UNKNOWN = "unknown"
)

// topology type
const (
	TYPE_POOL   = "pool"
	TYPE_SERVER = "server"
	TYPE_ZONE   = "zone"
)