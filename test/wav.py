import wave
import io

wave_bytes = open('audio.wav', 'rb').read()
with wave.open(io.BytesIO(wave_bytes), "rb") as wave_file:
    print("Sample width in bytes:", wave_file.getsampwidth())
    print("Sampling frequency:", wave_file.getframerate())
    print("Number of frames:", wave_file.getnframes())

wave_bytes = open('jfk.wav', 'rb').read()
with wave.open(io.BytesIO(wave_bytes), "rb") as wave_file:
    print("Sample width in bytes:", wave_file.getsampwidth())
    print("Sampling frequency:", wave_file.getframerate())
    print("Number of frames:", wave_file.getnframes())