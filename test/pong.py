import aiohttp
import asyncio
import librosa

async def main():
    async with aiohttp.ClientSession() as session:
        data = aiohttp.FormData()
        filename = 'jfk.wav'
        file = open(filename, 'rb')

        data.add_field('file',
                    file,
                    filename=filename)

        async with session.post('http://0.0.0.0:8080/transcribe', data=data) as resp:
            print(await resp.text())

asyncio.run(main())