package sentinel

import (
	"strconv"
	"testing"
	"time"

	"icode.baidu.com/baidu/gdp/got/ut"
)

func TestNew(t *testing.T) {
	type args struct {
		sentinelAddrs  []string
		sentinelName   string
		connectTimeout time.Duration
		readTimeout    time.Duration
		writeTimeout   time.Duration
	}

	type Want struct {
		Value  interface{}
		Assert func(actual interface{}, expected ...interface{}) string
	}

	utInst := ut.New(t, "Unit tests for FUNCTION: New ")
	defer utInst.RestoreAll()
	ag := utInst.NewAssertGroup()

	tests := []struct {
		testCaseTitle string
		args          args
		want          Want //sentinel
	}{
		// TODO: Add test cases.
	}

	for k, tt := range tests {
		got := New(tt.args.sentinelAddrs, tt.args.sentinelName, tt.args.connectTimeout, tt.args.readTimeout, tt.args.writeTimeout)
		ag.Add(strconv.Itoa(k)+" Test Case Of TestNew, Result Index:0 Value Compare", got, tt.want.Assert, tt.want.Value.(sentinel))
	}
	ag.Run()
}

func Test_sentinel_GetMasterAddrByName(t *testing.T) {
	type fields struct {
		Name           string
		Addrs          []string
		ConnectTimeout time.Duration
		ReadTimeout    time.Duration
		WriteTimeout   time.Duration
	}
	type args struct {
		masterName string
	}

	type Want struct {
		Value  interface{}
		Assert func(actual interface{}, expected ...interface{}) string
	}

	utInst := ut.New(t, "Unit tests for FUNCTION: GetMasterAddrByName ")
	defer utInst.RestoreAll()
	ag := utInst.NewAssertGroup()

	tests := []struct {
		testCaseTitle string
		fields        fields
		args          args
		want          Want //string
		wantErr       bool
	}{
		// TODO: Add test cases.
	}

	for k, tt := range tests {
		s := sentinel{
			Name:           tt.fields.Name,
			Addrs:          tt.fields.Addrs,
			ConnectTimeout: tt.fields.ConnectTimeout,
			ReadTimeout:    tt.fields.ReadTimeout,
			WriteTimeout:   tt.fields.WriteTimeout,
		}
		got, err := s.GetMasterAddrByName(tt.args.masterName)
		if err != nil {
			ag.Add(strconv.Itoa(k)+" Test Case Of Test_sentinel_GetMasterAddrByName, Error Value Compare", err != nil, ut.ShouldEqual, tt.wantErr)
			continue
		}
		ag.Add(strconv.Itoa(k)+" Test Case Of Test_sentinel_GetMasterAddrByName, Result Index:0 Value Compare", got, tt.want.Assert, tt.want.Value.(string))
	}
	ag.Run()
}

func Test_sentinel_Do(t *testing.T) {
	type fields struct {
		Name           string
		Addrs          []string
		ConnectTimeout time.Duration
		ReadTimeout    time.Duration
		WriteTimeout   time.Duration
	}
	type args struct {
		masterName string
		cmd        string
		option     []interface{}
	}

	type Want struct {
		Value  interface{}
		Assert func(actual interface{}, expected ...interface{}) string
	}

	utInst := ut.New(t, "Unit tests for FUNCTION: Do ")
	defer utInst.RestoreAll()
	ag := utInst.NewAssertGroup()

	tests := []struct {
		testCaseTitle string
		fields        fields
		args          args
		want          Want //interface{}
		wantErr       bool
	}{
		// TODO: Add test cases.
	}

	for k, tt := range tests {
		s := sentinel{
			Name:           tt.fields.Name,
			Addrs:          tt.fields.Addrs,
			ConnectTimeout: tt.fields.ConnectTimeout,
			ReadTimeout:    tt.fields.ReadTimeout,
			WriteTimeout:   tt.fields.WriteTimeout,
		}
		got, err := s.Do(tt.args.masterName, tt.args.cmd, tt.args.option...)
		if err != nil {
			ag.Add(strconv.Itoa(k)+" Test Case Of Test_sentinel_Do, Error Value Compare", err != nil, ut.ShouldEqual, tt.wantErr)
			continue
		}
		ag.Add(strconv.Itoa(k)+" Test Case Of Test_sentinel_Do, Result Index:0 Value Compare", got, tt.want.Assert, tt.want.Value.(interface{}))
	}
	ag.Run()
}
