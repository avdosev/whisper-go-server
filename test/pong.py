import aiohttp
import asyncio
import librosa
from ffmpeg.asyncio import FFmpeg
import io


async def main():
    async with aiohttp.ClientSession() as session:
        data = aiohttp.FormData()
        filename = 'audio.ogg'

        file = open(filename, 'rb')

        ffmpeg = (
            FFmpeg()
            .option("y")
            .input("pipe:0")
            .output(
                "puk.wav",
                {"codec:a": "pcm_s16le"},
                vn=None,
                ar='16000',
                f="wav",
            )
        )

        # audio = io.BytesIO(await ffmpeg.execute(file.read()))
        await ffmpeg.execute(file.read())

        # with open('puk.wav', 'wb') as f:
            # f.write(audio.getbuffer())
    
        data.add_field('file',
                    # audio.getbuffer(),
                    open('puk.wav', 'rb').read(),
                    filename='audio.wav')

        async with session.post('http://0.0.0.0:8080/transcribe', data=data) as resp:
            print(resp.status)
            print(await resp.text())

asyncio.run(main())