import hashlib

def authenticate(password: str) -> bool:
    encoded_password = password.encode()
    hashed_password = hashlib.sha256(encoded_password)

    return hashed_password.hexdigest() == "8f234f4c4584271ebfd4384564ad5324052848185180c7aec9413f2843e8cea5" # Password is anirudhisreallycool
