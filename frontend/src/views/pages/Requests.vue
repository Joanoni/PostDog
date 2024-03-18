<template>
    <div class="card flex flex-column">
        <div class="flex">
            <InputGroup>
                <Dropdown v-model="selectedMethod" :options="methods" optionLabel="name" placeholder="Select a City"
                    class="col-2" style="padding: 0px;" />
                <InputText type="text" v-model="url" class="col-8" />
                <Button label="Submit" class="col-2" @click="sendRequest" />
            </InputGroup>
        </div>
        <div class="divider-vertical"></div>
        <InputGroup>
            <div v-for="header of headers" :key="header.id">
                <InputGroup>
                    <InputText placeholder="Name" :value="header.name" class="col-5" />
                    <InputText placeholder="Value" :value="header.value" class="col-5" />
                    <InputGroupAddon class=" col-2">
                        <div class="flex flex-row">
                            <Checkbox v-model="header.active" :inputId="header.id" :binary="true" class="mr-1" />
                            <div class="flex align-items-center">
                                <label :for="header.id">Active</label>
                            </div>

                        </div>
                    </InputGroupAddon>
                </InputGroup>
            </div>
        </InputGroup>
        <div class="divider-vertical"></div>
        <FloatLabel>
            <Textarea v-model="inputBody" rows="15" style="width: 100%;" />
            <label>Body</label>
        </FloatLabel>
        <div class="divider-vertical"></div>
        <div class="flex flex-column">
            <div class="flex flex-row">
                <p>Status code:
                </p>
                <div v-if="response.statusCode" class="ml-1">
                    <p>{{ response.statusCode }}</p>
                </div>
            </div>
            <JsonViewer :value="response.jsonData" :expand-depth=5 copyable class="card">
            </JsonViewer>
        </div>
    </div>
</template>

<script>
import Backend from '../../service/Backend';

export default {
    data() {
        return {
            response: { jsonData: null, statusCode: null },
            inputBody: "",
            headers: [
                {
                    id: 0,
                    name: "Authentication",
                    value: "jwt",
                    active: true,
                },
                {
                    id: 0,
                    name: "Authentication",
                    value: "jwt",
                    active: true,
                },
            ],
            url: "http://localhost:7963/api/test/get",
            selectedMethod: { name: "GET" },
            methods: [
                { name: "GET", },
                { name: "POST", },
                { name: "PATCH", },
                { name: "DELETE", },
                { name: "CAIO", }
            ]
        };
    },
    mounted() {
        document.body.addEventListener('keydown', (event) => {
            if (event.key === "Enter" && (event.metaKey || event.ctrlKey)) {
                this.sendRequest()
            }
        });
    },
    methods: {
        sendRequest() {
            console.log("method: " + this.selectedMethod.name)
            console.log("url: " + this.url)
            console.log("inputBody: " + this.inputBody)
            let inputParsed = null
            if (this.inputBody != "" && this.inputBody != null && this.inputBody != undefined) {
                inputParsed = JSON.parse(this.inputBody)
            }
            Backend.request("post", "api/execute", { method: this.selectedMethod.name, url: this.url, payload: inputParsed }).then((resp) => {
                console.log(resp)
                this.response.jsonData = resp.data.data.body
                this.response.statusCode = resp.data.data.statusCode
            })
        }
    },
};
</script>

<style>
.divider-vertical {
    height: 20px;
}
</style>
