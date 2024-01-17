from typing import List

from hstest import CheckResult, dynamic_test, StageTest, TestCase, TestedProgram

FILENAME = 'forbidden_words.txt'
words = 'awful\natrocious\nharsh\ncrummy\ndreadful\nlousy'

T = 'True'
F = 'False'
MSG_ERROR_NOT_BAD = 'This word does not belong to bad words'
MSG_ERROR_BAD = 'This word belongs to bad words'


class CheckerStage2Test(StageTest):
    def da_test(self, word, flag, msg):
        pr = TestedProgram()
        pr.start()
        pr.execute(FILENAME)
        out = pr.execute(word).strip()
        if out.lower() != flag.lower():
            return CheckResult.wrong(msg)
        return CheckResult.correct()

    def generate(self) -> List[TestCase]:
        return [
            TestCase(dynamic_testing=i, files={FILENAME: words}) for i in [
                self.test1, self.test2, self.test3, self.test4, self.test5, self.test6,
                self.test7, self.test8, self.test9, self.test10, self.test11, self.test12,
            ]
        ]

    def test1(self):
        return self.da_test('awful', T, MSG_ERROR_BAD)

    def test2(self):
        return self.da_test('HARSH', T, MSG_ERROR_BAD)

    def test3(self):
        return self.da_test('awesome', F, MSG_ERROR_NOT_BAD)

    def test4(self):
        return self.da_test('LoUsE', F, MSG_ERROR_NOT_BAD)

    def test5(self):
        return self.da_test('crumble', F, MSG_ERROR_NOT_BAD)

    def test6(self):
        return self.da_test('dreadlocks', F, MSG_ERROR_NOT_BAD)

    def test7(self):
        return self.da_test('DREadFUL', T, MSG_ERROR_BAD)

    def test8(self):
        return self.da_test('crummy', T, MSG_ERROR_BAD)

    def test9(self):
        return self.da_test('atrium', F, MSG_ERROR_NOT_BAD)

    def test10(self):
        return self.da_test('lousy', T, MSG_ERROR_BAD)

    def test11(self):
        return self.da_test('atRocious', T, MSG_ERROR_BAD)

    def test12(self):
        return self.da_test('MARSH', F, MSG_ERROR_NOT_BAD)


if __name__ == '__main__':
    CheckerStage2Test().run_tests()
