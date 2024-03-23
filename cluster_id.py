import uuid
import base64

def generate_unique_base64_uuid():
    # Generate a UUID
    uuid_val = uuid.uuid4()
    # Convert UUID to bytes
    uuid_bytes = uuid_val.bytes
    # Encode bytes to Base64
    base64_uuid = base64.urlsafe_b64encode(uuid_bytes)
    # Convert bytes to string
    base64_uuid_str = base64_uuid.decode('utf-8')
    return base64_uuid_str

unique_base64_uuid = generate_unique_base64_uuid()
print("Unique Base64 UUID:", unique_base64_uuid)
