<center><h2><b>Week 11: 14.08 - 18.08.23</b></h2></center>

**P4 Custom Resource**

- This week I created P4 custom resource. This is the [commit](https://git.comnets.net/p4-work/p4-kube/-/commit/a71e7e002f48f088a4ad35543d7731ca60353799) for this change. Please review it.

- We can now create or delete the instances of P4 objects using manifest files.

- Next week, I will focus on creating custom controller to do magic with P4 resources.


**P4Runtime API for T4P4S**

- This week I spent a day on resolving issue with P4Runtime API. It still under progress.

    `./t4p4s.sh :l2fwd-gen model=v1model`

- Replace the file from t4p4s_mod to harware_indep/controlplane.c.py. Error comes `NameError: name 'hlir16' is not defined`. 

- Now replace all instances of hlir16 with hlir in controlplane.c.py. Error comes:
`AttributeError: Key 'match_type' not found in #486986@P4Table`.
After printing table, found out that table does not have *match_type* but *matchType*. Replace all the instances

- Error comes, `AttributeError: Key 'width' not found in #487032@KeyElement`
Changed width to size in function *get_key_byte_width*
Also chnaged *matchType_order* function a little bit.

- Error comes; 
```
[RUN CONTROLLER] dpdk_l2fwd_controller (default for l2fwd-gen@std)
[COMPILE  P4-16] ./examples/l2fwd-gen.p4 @std
table actions
 0 #485632#ActionListElement[expression.method.path.name=forward, action_object.5, annotations**0]
1 #485638#ActionListElement[expression.method.path.name=bcast, action_object.5, annotations**0]
2 #485644#ActionListElement[expression.method.path.name=NoAction_1, action_object.5, annotations**1]
table actions
 0 #485697#ActionListElement[expression.method.path.name=mac_learn, action_object.5, annotations**0]
1 #485703#ActionListElement[expression.method.path.name=_nop, action_object.5, annotations**0]
2 #485709#ActionListElement[expression.method.path.name=NoAction_2, action_object.5, annotations**1]
[COMPILE SWITCH]
[47/58] Compiling C object l2fwd-gen.p/srcgen_controlplane.c.o
FAILED: l2fwd-gen.p/srcgen_controlplane.c.o 
ccache clang-10 -Il2fwd-gen.p -I. -I.. -I../srcgen -I../srcgen/multi -I../../../src/testing/includes -I../../../src/hardware_dep/dpdk/includes -I../../../src/hardware_dep/shared/includes -I../../../src/hardware_dep/shared/ctrl_plane -I/usr/local/include -I/usr/include/libnl3 -fvisibility=hidden -fcolor-diagnostics -D_FILE_OFFSET_BITS=64 -O2 -g -pthread -include rte_config.h -march=native -Wno-parentheses-equality -D__TARGET_V1__ -MD -MQ l2fwd-gen.p/srcgen_controlplane.c.o -MF l2fwd-gen.p/srcgen_controlplane.c.o.d -o l2fwd-gen.p/srcgen_controlplane.c.o -c ../srcgen/controlplane.c
../srcgen/controlplane.c:9:10: fatal error: 'p4rt/device_mgr.h' file not found
#include "p4rt/device_mgr.h"
         ^~~~~~~~~~~~~~~~~~~
1 error generated.
[56/58] Compiling C object l2fwd-gen.p/srcgen_multi_dataplane.stage_3.c.o
ninja: build stopped: subcommand failed.

Error: C compilation using ninja failed, see log build/l2fwd-gen@std-v1model/log/22_ninja.txt (error code: 1)

```

- Don't yet know how to resolve this error. I think it must have something to do with makefiles.
