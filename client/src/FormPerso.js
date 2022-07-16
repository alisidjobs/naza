import React, { Component } from 'react'
import {
    Button,
    Checkbox,
    Form,
    Input,
    Radio,
    Select,
    TextArea,
} from 'semantic-ui-react'

const options = [
    { key: 'm', text: 'Male', value: 'male' },
    { key: 'f', text: 'Female', value: 'female' },
]

class FormPerso extends Component {
    state = {}

    handleChange = (e, { value }) => this.setState({ value })

    render() {
        const { value } = this.state
        return (
            <div>
            <Form>
                <Form.Group inline>
                    <Form.Field
                        width={4}
                        control={Input}
                        label='Name'
                        placeholder='Name'
                    />
                        <Form.Field
                            control={Radio}
                            label='Male'
                            value='1'
                            checked={value === '1'}
                            onChange={this.handleChange}
                        />
                        <Form.Field
                            control={Radio}
                            label='Female'
                            value='2'
                            checked={value === '2'}
                            onChange={this.handleChange}
                        />
                </Form.Group>
                <Form.Field
                    width={4}
                    control={Select}
                    label='Race'
                    options={options}
                    placeholder='Race'
                />
                <Form.Field
                    width={4}
                    control={Select}
                    label='Class'
                    options={options}
                    placeholder='Class'
                />
                <Form.Field control={Button}>Submit</Form.Field>
            </Form>
                <div>
                    <pre>{JSON.stringify({ value }, null, 2)}</pre></div>
            </div>
        )
    }
}

export default FormPerso