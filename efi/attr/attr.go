package attr

import (
	"errors"
	"os"
)

/* The code below won't work correctly if this tool is built
 * for a 64-bit big endian platform.
 * See https://github.com/golang/go/issues/45585 for context. */

const (
	// from /usr/include/linux/fs.h
	FS_SECRM_FL        = 0x00000001 /* Secure deletion */
	FS_UNRM_FL         = 0x00000002 /* Undelete */
	FS_COMPR_FL        = 0x00000004 /* Compress file */
	FS_SYNC_FL         = 0x00000008 /* Synchronous updates */
	FS_IMMUTABLE_FL    = 0x00000010 /* Immutable file */
	FS_APPEND_FL       = 0x00000020 /* writes to file may only append */
	FS_NODUMP_FL       = 0x00000040 /* do not dump file */
	FS_NOATIME_FL      = 0x00000080 /* do not update atime */
	FS_DIRTY_FL        = 0x00000100
	FS_COMPRBLK_FL     = 0x00000200 /* One or more compressed clusters */
	FS_NOCOMP_FL       = 0x00000400 /* Don't compress */
	FS_ECOMPR_FL       = 0x00000800 /* Compression error */
	FS_BTREE_FL        = 0x00001000 /* btree format dir */
	FS_INDEX_FL        = 0x00001000 /* hash-indexed directory */
	FS_IMAGIC_FL       = 0x00002000 /* AFS directory */
	FS_JOURNAL_DATA_FL = 0x00004000 /* Reserved for ext3 */
	FS_NOTAIL_FL       = 0x00008000 /* file tail should not be merged */
	FS_DIRSYNC_FL      = 0x00010000 /* dirsync behaviour (directories only) */
	FS_TOPDIR_FL       = 0x00020000 /* Top of directory hierarchies*/
	FS_EXTENT_FL       = 0x00080000 /* Extents */
	FS_DIRECTIO_FL     = 0x00100000 /* Use direct i/o */
	FS_NOCOW_FL        = 0x00800000 /* Do not cow file */
	FS_PROJINHERIT_FL  = 0x20000000 /* Create with parents projid */
	FS_RESERVED_FL     = 0x80000000 /* reserved for ext2 lib */
)

var ErrIsImmutable = errors.New("file is immutable")

// checks if the file is immutable
func IsImmutable(p string) error {
	a, err := GetAttr(p)
	if err != nil {
		return err
	}
	if (a & FS_IMMUTABLE_FL) != 0 {
		return ErrIsImmutable
	}
	return nil
}

func UnsetImmutable(p string) error {
	a, err := GetAttr(p)
	if err != nil {
		return err
	}
	a &= ^FS_IMMUTABLE_FL
	if err := SetAttr(p, a); err != nil {
		return err
	}
	return nil
}

// GetAttr retrieves the attributes of a file on a linux filesystem
func GetAttr(path string) (int32, error) {
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return GetAttrFromFile(f)
}

// SetAttr sets the attributes of a file on a linux filesystem to the given value
func SetAttr(path string, attr int32) error {
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	return SetAttrOnFile(f, attr)
}
