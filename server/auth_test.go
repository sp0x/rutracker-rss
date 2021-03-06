package server

import (
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"

	"github.com/sp0x/torrentd/config/mocks"
)

func TestServer_checkAPIKey(t *testing.T) {
	g := NewGomegaWithT(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config := mocks.NewMockConfig(ctrl)
	config.EXPECT().GetInt("port").Return(3333).Times(1)
	config.EXPECT().GetString("hostname").Return("").Times(1)
	config.EXPECT().GetBytes("api_key").Return(nil).Times(1)
	// Test
	s := NewServer(config)
	g.Expect(s.checkAPIKey("")).Should(BeFalse())

	s.Params.APIKey = []byte("demokey")
	ok := s.checkAPIKey("demokey")
	g.Expect(ok).Should(BeTrue())

	s.Params.APIKey = nil
	s.Params.Passphrase = "serverpass"
	ok = s.checkAPIKey("serverpass")
	g.Expect(ok).Should(BeFalse())

	ok = s.checkAPIKey("cd2234c6b7755b8dd230bdbc84544c38")
	g.Expect(ok).Should(BeTrue())
}

func TestServer_sharedKey(t *testing.T) {
	g := NewGomegaWithT(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config := mocks.NewMockConfig(ctrl)
	config.EXPECT().GetInt("port").Return(3333).Times(1)
	config.EXPECT().GetString("hostname").Return("").Times(1)
	config.EXPECT().GetBytes("api_key").Return(nil).Times(1)
	// Test
	s := NewServer(config)
	bytes := s.sharedKey()
	bytes2 := s.sharedKey()

	g.Expect(len(bytes)).Should(Equal(32))
	g.Expect(bytes).ShouldNot(Equal(bytes2))

	s.Params.APIKey = []byte("demokey")
	bytes = s.sharedKey()
	g.Expect(bytes).Should(Equal([]byte("demokey")))

	s.Params.APIKey = nil
	s.Params.Passphrase = "serverpass"
	bytes = s.sharedKey()
	g.Expect(len(bytes)).Should(Equal(32))
	g.Expect(bytes).Should(Equal([]byte("cd2234c6b7755b8dd230bdbc84544c38")))
}
