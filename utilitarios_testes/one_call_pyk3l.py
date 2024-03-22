import sys
import logging
logging.basicConfig(format='%(asctime)s.%(msecs)03d %(message)s', datefmt='%m/%d/%Y %H:%M:%S')
import time
k3l = pyk3l.k3lclient.K3L(3, 3, 0)
sip = k3l.device(99)
sip[0].send_command(pyk3l.defines.KCommand.CM_ENABLE_HMP_ANALYTICS)
sip[0].send_command(pyk3l.defines.KCommand.CM_ENABLE_HMP_ANALYTICS)
sip[1].send_command(pyk3l.defines.KCommand.CM_ENABLE_HMP_ANALYTICS)
sip[1].send_command(pyk3l.defines.KCommand.CM_MAKE_CALL, orig_addr=1234, dest_addr=12332, network_dest_addr="127.0.0.1")
sip[0].send_command(pyk3l.defines.KCommand.CM_RINGBACK)  # 180
sip[0].send_command(pyk3l.defines.KCommand.CM_PRE_CONNECT)  # 183
sip[1].send_command(pyk3l.defines.KCommand.CM_PLAY_FROM_FILE, 'G:/myfiles/REC_2022_05_24_22_26_03_6411993907180_eb2605a883fe430e85527472277990d0-pcm.wav')
time.sleep(10)
sip[0].send_command(pyk3l.defines.KCommand.CM_CONNECT)
time.sleep(10)
sip[0].send_command(pyk3l.defines.KCommand.CM_DISCONNECT)
time.sleep(0)