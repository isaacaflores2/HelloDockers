FROM mcr.microsoft.com/dotnet/sdk:3.1 as build

WORKDIR /app

ADD hello.txt .

COPY ./csharp .

RUN dotnet restore

RUN dotnet publish -c release -o ./build

FROM mcr.microsoft.com/dotnet/aspnet:3.1-alpine

WORKDIR /app

COPY --from=build /app/hello.txt .

COPY --from=build /app/bin/release/netcoreapp3.1/ .

ENTRYPOINT ["dotnet", "hello.dll", "hello.txt"]