from hstest import CheckResult, StageTest, TestCase

MSG_ERROR_CONTENT = 'Incorrect reading from the file'


class CheckerStage1Test(StageTest):

    words = 'awful\natrocious\nharsh\ncrummy\ndreadful\nlousy'

    def generate(self):
        filename = 'forbidden_words.txt'
        return [
            TestCase(stdin=filename, files={
                filename: self.words
            }, check_function=self.check),
            TestCase(files={
                filename: '\n'
            }, stdin=filename, check_function=self.check_empty),
            TestCase(files={
                'words.txt': self.words
            }, stdin='words.txt', check_function=self.check),
        ]

    def check(self, reply: str, attach) -> CheckResult:
        words = self.words.split('\n')
        rp_split = reply.split()
        if len(words) != len(rp_split):
            return CheckResult.wrong(MSG_ERROR_CONTENT)
        for i, word in enumerate(rp_split):
            if word != words[i]:
                return CheckResult.wrong(MSG_ERROR_CONTENT)
        return CheckResult.correct()

    def check_empty(self, reply: str, attach) -> CheckResult:
        if reply.strip('\n') != '':
            return CheckResult.wrong(MSG_ERROR_CONTENT)
        return CheckResult.correct()


if __name__ == '__main__':
    CheckerStage1Test().run_tests()
