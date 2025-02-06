import unittest
from main import *

class AuthenticationTest(unittest.TestCase):

    def test_correct(self):
        self.assertTrue(authenticate("anirudhisreallycool"))
    
    def test_incorrect(self):
        self.assertFalse(authenticate("anirudhisnotreallycool"))

    def test_empty(self):
        self.assertFalse(authenticate(""))