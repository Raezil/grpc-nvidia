from openai import OpenAI
import os
import sys
import asyncio

class NVIDIA:
    def __init__(self, model) -> None:
        api_key = os.environ.get("NVIDIA_API")
        self.model = model
        self.client = OpenAI(
            base_url = "https://integrate.api.nvidia.com/v1",
            api_key = api_key
        )
    
    def call(self, prompt):
        completion = self.client.chat.completions.create(
            model=self.model,
            messages=[{"role":"user","content":prompt}],
            temperature=0.2,
            top_p=0.7,
            max_tokens=1024,
            stream=True
        )

        for chunk in completion:
            if chunk.choices[0].delta.content is not None:
                print(chunk.choices[0].delta.content, end="")

async def main():
    args = sys.argv
    print(args)
    nvidia = NVIDIA(model = args[2])
    nvidia.call(args[1])

if __name__ == "__main__":
    asyncio.run(main())
