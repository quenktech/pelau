package pelau_test

import (
	"code.google.com/p/gomock/gomock"
	"github.com/metasansana/pelau"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/thirdparty/gomocktestreporter"
)

var _ = Describe("BridgeRouter", func() {

	var (
		bridge  pelau.Router
		router  *pelau.MockRouter
		route   pelau.Route
		mockCtl *gomock.Controller
	)

	BeforeEach(func() {

		mockCtl = gomock.NewController(gomocktestreporter.New())
		router = pelau.NewMockRouter(mockCtl)
		route = pelau.NewMockRoute(mockCtl)
		bridge = pelau.Cook(router)

	})

	AfterEach(func() {

		mockCtl.Finish()

	})

	It("should bridge Get calls", func() {

		router.EXPECT().Get(route)
		bridge.Get(route)

	})
	It("should bridge Post calls", func() {

		router.EXPECT().Post(route)
		bridge.Post(route)

	})
	It("should bridge Put calls", func() {

		router.EXPECT().Put(route)
		bridge.Put(route)

	})
	It("should bridge Delete calls", func() {

		router.EXPECT().Delete(route)
		bridge.Delete(route)

	})
	It("should bridge Head calls", func() {

		router.EXPECT().Head(route)
		bridge.Head(route)

	})

	It("should bridge Use calls", func() {

		router.EXPECT().Use(route)
		bridge.Use(route)

	})

	It("should bridge Bind calls", func() {

		cb := func() {}
		router.EXPECT().Bind("addr", cb)
		bridge.Bind("addr", cb)

	})

	It("should bridge ServerHttp calls", func() {

	})

})
