package hashtable

import "testing"

import (
    "os"
    "math/rand"
)

import bs "file-structures/block/byteslice"

func init() {
    if urandom, err := os.Open("/dev/urandom"); err != nil {
        return
    } else {
        seed := make([]byte, 8)
        if _, err := urandom.Read(seed); err == nil {
            rand.Seed(int64(bs.ByteSlice(seed).Int64()))
        }
    }
}

func randstr(length int) String {
    if urandom, err := os.Open("/dev/urandom"); err != nil {
        panic(err)
    } else {
        slice := make([]byte, length)
        if _, err := urandom.Read(slice); err != nil {
            panic(err)
        }
        urandom.Close()
        return String(slice)
    }
    panic("unreachable")
}

func TestMake(t *testing.T) {
    NewHashTable(12)
    t.Log(String("asdf").Hash())
}

func TestHashable(t *testing.T) {
    a := String("asdf")
    b := String("asdf")
    c := String("csfd")
    if !a.Equals(b) {
        t.Error("a != b")
    }
    if a.Hash() != b.Hash() {
        t.Error("hash(a) != hash(b)")
    }
    if a.Equals(c) {
        t.Error("a == c")
    }
    if a.Hash() != c.Hash() {
        t.Error("hash(a) != hash(c)")
    }
}

func TestPutHasGetRemove(t *testing.T) {

    type record struct {
        key String
        value String
    }

    records := make([]*record, 100)
    table := NewHashTable(100)

    ranrec := func() *record {
        return &record{ randstr(20), randstr(20) }
    }

    for i := range records {
        r := ranrec()
        records[i] = r
        err := table.Put(r.key, String(""))
        if err != nil {
            t.Error(err)
        }
        err = table.Put(r.key, r.value)
        if err != nil {
            t.Error(err)
        }
        if (table.(*hash)).size != (i+1) {
            t.Error("size was wrong", (table.(*hash)).size, i+1)
        }
    }

    for _, r := range records {
        if has := table.Has(r.key); !has {
            t.Error("Missing key")
        }
        if has := table.Has(randstr(12)); has {
            t.Error("Table has extra key")
        }
        if val, err := table.Get(r.key); err != nil {
            t.Error(err)
        } else if !(val.(String)).Equals(r.value) {
            t.Error("wrong value")
        }
    }

    for i, x := range records {
        if val, err := table.Remove(x.key); err != nil {
            t.Error(err)
        } else if !(val.(String)).Equals(x.value) {
            t.Error("wrong value")
        }
        for _, r := range records[i+1:] {
            if has := table.Has(r.key); !has {
                t.Error("Missing key")
            }
            if has := table.Has(randstr(12)); has {
                t.Error("Table has extra key")
            }
            if val, err := table.Get(r.key); err != nil {
                t.Error(err)
            } else if !(val.(String)).Equals(r.value) {
                t.Error("wrong value")
            }
        }
        if (table.(*hash)).size != (len(records) - (i+1)) {
            t.Error("size was wrong", (table.(*hash)).size, (len(records) - (i+1)))
        }
    }
}


