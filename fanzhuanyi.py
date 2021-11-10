# -*- coding: utf-8 -*-
co = {
    "d": 3.2,
    "m": 5.13,
    "s": 1.56,
    "h": "E"
}

print("%5.2f° %5.2f\' %5.2f\" \n%s" %(co["d"], co["m"], co["s"], co["h"]))

print(f"{co['d']}° {co['m']}\' {co['s']}\" \n{co['h']}")