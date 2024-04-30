import uuid


def getmac():
    mac = uuid.getnode()
    mac_address = ":".join(
        ["{:02x}".format((mac >> elements) & 0xFF) for elements in range(0, 2 * 6, 2)][
            ::-1
        ]
    )
    return mac_address


print("MAC Address:", getmac())
