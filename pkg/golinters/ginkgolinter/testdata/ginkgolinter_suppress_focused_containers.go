//golangcitest:args --disable-all -Eginkgolinter
//golangcitest:config_path testdata/ginkgolinter_suppress_focused_containers.yml
package ginkgolinter

import (
	"errors"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func LenUsecase_focus() {
	var fakeVarUnderTest []int
	Expect(fakeVarUnderTest).Should(BeEmpty())     // valid
	Expect(fakeVarUnderTest).ShouldNot(HaveLen(5)) // valid

	Expect(len(fakeVarUnderTest)).Should(Equal(0))           // want "ginkgo-linter: wrong length assertion. Consider using .Expect\\(fakeVarUnderTest\\)\\.Should\\(BeEmpty\\(\\)\\)"
	Expect(len(fakeVarUnderTest)).ShouldNot(Equal(2))        // want "ginkgo-linter: wrong length assertion. Consider using .Expect\\(fakeVarUnderTest\\)\\.ShouldNot\\(HaveLen\\(2\\)\\)"
	Expect(len(fakeVarUnderTest)).To(BeNumerically("==", 0)) // // want "ginkgo-linter: wrong length assertion. Consider using .Expect\\(fakeVarUnderTest\\)\\.To\\(BeEmpty\\(\\)\\)"

	fakeVarUnderTest = append(fakeVarUnderTest, 3)
	Expect(len(fakeVarUnderTest)).ShouldNot(Equal(0))        // want "ginkgo-linter: wrong length assertion. Consider using .Expect\\(fakeVarUnderTest\\)\\.ShouldNot\\(BeEmpty\\(\\)\\)"
	Expect(len(fakeVarUnderTest)).Should(Equal(1))           // want "ginkgo-linter: wrong length assertion. Consider using .Expect\\(fakeVarUnderTest\\)\\.Should\\(HaveLen\\(1\\)\\)"
	Expect(len(fakeVarUnderTest)).To(BeNumerically(">", 0))  // want "ginkgo-linter: wrong length assertion. Consider using .Expect\\(fakeVarUnderTest\\)\\.ToNot\\(BeEmpty\\(\\)\\)"
	Expect(len(fakeVarUnderTest)).To(BeNumerically(">=", 1)) // want "ginkgo-linter: wrong length assertion. Consider using .Expect\\(fakeVarUnderTest\\)\\.ToNot\\(BeEmpty\\(\\)\\)"
	Expect(len(fakeVarUnderTest)).To(BeNumerically("!=", 0)) // want "ginkgo-linter: wrong length assertion. Consider using .Expect\\(fakeVarUnderTest\\)\\.ToNot\\(BeEmpty\\(\\)\\)"
}

func NilUsecase_focus() {
	y := 5
	x := &y
	Expect(x == nil).To(Equal(true)) // want "ginkgo-linter: wrong nil assertion. Consider using .Expect\\(x\\)\\.To\\(BeNil\\(\\)\\)"
	Expect(nil == x).To(Equal(true)) // want "ginkgo-linter: wrong nil assertion. Consider using .Expect\\(x\\)\\.To\\(BeNil\\(\\)\\)"
	Expect(x != nil).To(Equal(true)) // want "ginkgo-linter: wrong nil assertion. Consider using .Expect\\(x\\)\\.ToNot\\(BeNil\\(\\)\\)"
	Expect(x == nil).To(BeTrue())    // want "ginkgo-linter: wrong nil assertion. Consider using .Expect\\(x\\)\\.To\\(BeNil\\(\\)\\)"
	Expect(x == nil).To(BeFalse())   // want "ginkgo-linter: wrong nil assertion. Consider using .Expect\\(x\\)\\.ToNot\\(BeNil\\(\\)\\)"
}
func BooleanUsecase_focus() {
	x := true
	Expect(x).To(Equal(true)) // want "ginkgo-linter: wrong boolean assertion. Consider using .Expect\\(x\\)\\.To\\(BeTrue\\(\\)\\)"
	x = false
	Expect(x).To(Equal(false)) // want "ginkgo-linter: wrong boolean assertion. Consider using .Expect\\(x\\)\\.To\\(BeFalse\\(\\)\\)"
}

func ErrorUsecase_focus() {
	err := errors.New("fake error")
	funcReturnsErr := func() error { return err }

	Expect(err).To(BeNil())              // want "ginkgo-linter: wrong error assertion. Consider using .Expect\\(err\\)\\.ToNot\\(HaveOccurred\\(\\)\\)"
	Expect(err == nil).To(Equal(true))   // want "ginkgo-linter: wrong error assertion. Consider using .Expect\\(err\\)\\.ToNot\\(HaveOccurred\\(\\)\\)"
	Expect(err == nil).To(BeFalse())     // want "ginkgo-linter: wrong error assertion. Consider using .Expect\\(err\\)\\.To\\(HaveOccurred\\(\\)\\)"
	Expect(err != nil).To(BeTrue())      // want "ginkgo-linter: wrong error assertion. Consider using .Expect\\(err\\)\\.To\\(HaveOccurred\\(\\)\\)"
	Expect(funcReturnsErr()).To(BeNil()) // want "ginkgo-linter: wrong error assertion. Consider using .Expect\\(funcReturnsErr\\(\\)\\)\\.To\\(Succeed\\(\\)\\)"
}

func HaveLen0Usecase_focus() {
	x := make([]string, 0)
	Expect(x).To(HaveLen(0)) // want "ginkgo-linter: wrong length assertion. Consider using .Expect\\(x\\)\\.To\\(BeEmpty\\(\\)\\)"
}

func WrongComparisonUsecase_focus() {
	x := 8
	Expect(x == 8).To(BeTrue())    // want "ginkgo-linter: wrong comparison assertion. Consider using .Expect\\(x\\)\\.To\\(Equal\\(8\\)\\)"
	Expect(x < 9).To(BeTrue())     // want "ginkgo-linter: wrong comparison assertion. Consider using .Expect\\(x\\)\\.To\\(BeNumerically\\(\"<\", 9\\)\\)"
	Expect(x < 7).To(Equal(false)) // want "ginkgo-linter: wrong comparison assertion. Consider using .Expect\\(x\\)\\.ToNot\\(BeNumerically\\(\"<\", 7\\)\\)"

	p1, p2 := &x, &x
	Expect(p1 == p2).To(Equal(true)) // want "ginkgo-linter: wrong comparison assertion. Consider using .Expect\\(p1\\).To\\(BeIdenticalTo\\(p2\\)\\)"
}

func slowInt_focus() int {
	time.Sleep(time.Second)
	return 42
}

func WrongEventuallyWithFunction_focus() {
	Eventually(slowInt_focus).Should(Equal(42))   // valid
	Eventually(slowInt_focus()).Should(Equal(42)) // want "ginkgo-linter: use a function call in Eventually. This actually checks nothing, because Eventually receives the function returned value, instead of function itself, and this value is never changed. Consider using .Eventually\\(slowInt_focus\\)\\.Should\\(Equal\\(42\\)\\)"
}

var _ = Describe("Should warn for focused containers", func() {
	FContext("should not allow FContext", func() { // want "ginkgo-linter: Focus container found. This is used only for local debug and should not be part of the actual source code. Consider to replace with \"Context\""
		FWhen("should not allow FWhen", func() { // want "ginkgo-linter: Focus container found. This is used only for local debug and should not be part of the actual source code. Consider to replace with \"When\""
			FIt("should not allow FIt", func() { // want "ginkgo-linter: Focus container found. This is used only for local debug and should not be part of the actual source code. Consider to replace with \"It\""

			})
		})
	})
})
