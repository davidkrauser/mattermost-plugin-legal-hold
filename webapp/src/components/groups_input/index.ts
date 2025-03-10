import {connect} from 'react-redux';
import {AnyAction, bindActionCreators, Dispatch} from 'redux';

import GroupsInput from './groups_input.jsx';

// Function to search groups via the plugin API
const searchGroups = (term: string) => {
    return async () => {
        try {
            const response = await fetch(`/plugins/com.mattermost.plugin-legal-hold/api/v1/groups/search?prefix=${encodeURIComponent(term)}`);
            if (!response.ok) {
                throw new Error('Failed to search groups');
            }
            return await response.json();
        } catch (error) {
            console.error('Error searching groups:', error);
            return [];
        }
    };
};

function mapDispatchToProps(dispatch: Dispatch<AnyAction>) {
    return {
        actions: bindActionCreators({
            searchGroups,
        }, dispatch),
    };
}

export default connect(null, mapDispatchToProps)(GroupsInput);
